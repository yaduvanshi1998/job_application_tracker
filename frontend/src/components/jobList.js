import { useEffect, useState } from "react";
import { getJobs, deleteJob, updateJob } from "../api/jobApi";

function JobList() {
  const [jobs, setJobs] = useState([]);

  useEffect(() => {
    fetchJobs();
  }, []);

  const fetchJobs = async () => {
    try {
      const data = await getJobs();
      setJobs(data);
    } catch (error) {
      console.error("Failed to fetch jobs:", error);
    }
  };

  const handleStatusChange = (id, newStatus) => {
    setJobs((prevJobs) =>
      prevJobs.map((job) =>
        job.id === id ? { ...job, status: newStatus } : job
      )
    );
  };

  const handleUpdate = async (id) => {
    try {
      const job = jobs.find((j) => j.id === id);
      if (!job) return;

      const updatedJob = {
        company: job.company,
        role: job.role,
        status: job.status,
        applied_date: job.applied_date,
        interview_date: job.interview_date,
      };

      const success = await updateJob(id, updatedJob);

      if (!success) {
        // Update failed — re-fetch to revert local state to what the backend has
        await fetchJobs();
      }
      // On success, local state is already correct so no re-fetch needed
    } catch (error) {
      console.error("Failed to update job:", error);
      await fetchJobs(); // Revert on unexpected error
    }
  };

  const handleDelete = async (id) => {
    try {
      await deleteJob(id);
      await fetchJobs();
    } catch (error) {
      console.error("Failed to delete job:", error);
    }
  };

  return (
    <div>
      <h2>Job List</h2>
      {jobs.length === 0 && <p>No jobs added yet.</p>}

      {jobs.map((job) => (
        <div
          key={job.id}
          style={{
            marginBottom: "10px",
            borderBottom: "1px solid #ccc",
            paddingBottom: "5px",
          }}
        >
          <strong>{job.company}</strong> | {job.role} |{" "}
          <select
            value={job.status}
            onChange={(e) => handleStatusChange(job.id, e.target.value)}
          >
            <option value="Applied">Applied</option>
            <option value="Interview">Interview</option>
            <option value="Offer">Offer</option>
            <option value="Rejected">Rejected</option>
          </select>{" "}
          <button onClick={() => handleUpdate(job.id)}>Update</button>{" "}
          <button onClick={() => handleDelete(job.id)}>Delete</button>
        </div>
      ))}
    </div>
  );
}

export default JobList;