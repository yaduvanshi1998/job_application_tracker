const API_URL = "http://localhost:8080";

export const getJobs = async () => {
  try {
    const res = await fetch(`${API_URL}/jobs`);
    if (!res.ok) throw new Error("Failed to fetch jobs");
    const data = await res.json();

    return data.map((job) => ({
      ...job,
      applied_date: job.applied_date || new Date().toISOString(),
      interview_date: job.interview_date || null,
    }));
  } catch (err) {
    console.error("Error fetching jobs:", err);
    return [];
  }
};

export const addJob = async (job) => {
  try {
    const res = await fetch(`${API_URL}/jobs`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(job),
    });
    if (!res.ok) throw new Error("Failed to add job");
    return await res.json();
  } catch (err) {
    console.error("Error adding job:", err);
    return null;
  }
};

export const updateJob = async (id, job) => {
  try {
    const res = await fetch(`${API_URL}/jobs/${id}`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(job),
    });
    if (!res.ok) throw new Error("Failed to update job");
    return true;
  } catch (err) {
    console.error("Error updating job:", err);
    return false;
  }
};

export const deleteJob = async (id) => {
  try {
    const res = await fetch(`${API_URL}/jobs/${id}`, { method: "DELETE" });
    if (!res.ok) throw new Error("Failed to delete job");
    return true;
  } catch (err) {
    console.error("Error deleting job:", err);
    return false;
  }
};
