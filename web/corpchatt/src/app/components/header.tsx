import React from "react";
import Link from "next/link";

export function Header() {
  return (
    <header style={{
      padding: "1rem 2rem",
      backgroundColor: "#1f2937",
      color: "white",
      fontWeight: "bold",
      fontSize: "1.5rem",
      display: "flex",
      alignItems: "center",
      justifyContent: "space-between",
    }}>
      <div>My App Logo</div>
      <nav>
        <Link href="/" style={{ color: "white", marginRight: "1rem", textDecoration: "none" }}>
          Home
        </Link>
        <Link href="/about" style={{ color: "white", textDecoration: "none" }}>
          About
        </Link>
      </nav>
    </header>
  );
}
