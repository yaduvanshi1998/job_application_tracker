import React from "react";
import NotificationBell from "./components/NotificationBell";
import JobForm from "./components/JobForm";
import JobList from "./components/JobList";

function App() {
  return (
    <div style={{ padding: "20px" }}>
      <h1>Job Application Tracker</h1>
      <NotificationBell />
      <JobForm />
      <hr />
      <JobList />
    </div>
  );
}

export default App;