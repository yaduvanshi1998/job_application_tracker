import React from "react";
import JobForm from "./components/JobForm";
import JobList from "./components/JobList";

function App() {
  return (
    <div style={{ padding: "20px" }}>
      <h1>Job Application Tracker</h1>
      <JobForm />
      <hr />
      <JobList />
    </div>
  );
}

export default App;