import { useEffect, useState } from "react";

const API_URL = "http://localhost:8081";

function NotificationBell() {
  const [notifications, setNotifications] = useState([]);
  const [open, setOpen] = useState(false);

  useEffect(() => {
    const poll = async () => {
      const res = await fetch(`${API_URL}/notifications`);
      const data = await res.json();
      setNotifications(data);
    };

    poll();
    const interval = setInterval(poll, 1000);
    return () => clearInterval(interval);
  }, []);

  const dismiss = async (id) => {
    await fetch(`${API_URL}/notifications/${id}`, { method: "DELETE" });
    setNotifications((prev) => prev.filter((n) => n.id !== id));
  };

  return (
    <div
      style={{
        width: "100%",
        display: "flex",
        justifyContent: "flex-end",
        position: "relative",
        zIndex: 1000,
      }}
    >
      <div style={{ position: "relative", display: "inline-block" }}>
        <button
          onClick={() => setOpen((o) => !o)}
          style={{
            fontSize: "22px",
            background: "none",
            border: "none",
            cursor: "pointer",
          }}
        >
          🔔
          {notifications.length > 0 && (
            <span
              style={{
                position: "absolute",
                top: 0,
                right: 0,
                background: "red",
                color: "white",
                borderRadius: "50%",
                fontSize: "11px",
                padding: "1px 5px",
              }}
            >
              {notifications.length}
            </span>
          )}
        </button>

        {open && (
          <div
            style={{
              position: "absolute",
              right: 0,
              top: "36px",
              background: "white",
              border: "1px solid #ccc",
              borderRadius: "8px",
              width: "280px",
              boxShadow: "0 4px 12px rgba(0,0,0,0.1)",
              zIndex: 9999,
            }}
          >
            {notifications.length === 0 ? (
              <p style={{ padding: "12px", margin: 0, color: "#888" }}>
                No notifications
              </p>
            ) : (
              notifications.map((n) => (
                <div
                  key={n.id}
                  style={{
                    display: "flex",
                    justifyContent: "space-between",
                    alignItems: "center",
                    padding: "10px 12px",
                    borderBottom: "1px solid #eee",
                  }}
                >
                  <span style={{ fontSize: "13px", color: "#333" }}>
                    {n.message}
                  </span>
                  <button
                    onClick={() => dismiss(n.id)}
                    style={{
                      background: "none",
                      border: "none",
                      cursor: "pointer",
                      color: "#999",
                      fontSize: "16px",
                    }}
                  >
                    ×
                  </button>
                </div>
              ))
            )}
          </div>
        )}
      </div>
    </div>
  );
}

export default NotificationBell;
