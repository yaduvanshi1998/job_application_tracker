import { useState } from "react";
import { addJob } from "../api/jobApi";

function JobForm({ onJobAdded }) {
  const [company, setCompany] = useState("");
  const [role, setRole] = useState("");
  const [status, setStatus] = useState("Applied");

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      await addJob({ company, role, status });
      setCompany("");
      setRole("");
      setStatus("Applied");

      if (onJobAdded) onJobAdded();
    } catch (error) {
      console.error("Failed to add job:", error);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        value={company}
        onChange={(e) => setCompany(e.target.value)}
        placeholder="Company"
      />
      <input
        value={role}
        onChange={(e) => setRole(e.target.value)}
        placeholder="Role"
      />
      <select value={status} onChange={(e) => setStatus(e.target.value)}>
        <option value="Applied">Applied</option>
        <option value="Interview">Interview</option>
        <option value="Offer">Offer</option>
        <option value="Rejected">Rejected</option>
      </select>
      <button type="submit">Add Job</button>
    </form>
  );
}

export default JobForm;