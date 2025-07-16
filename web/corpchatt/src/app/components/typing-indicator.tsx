"use client";
import { useEffect, useState } from "react";
import { socket } from "@/lib/socket";

export const TypingIndicator = () => {
  const [typingUsers, setTypingUsers] = useState<string[]>([]);

  useEffect(() => {
    socket.on("typing:start", (name) =>
      setTypingUsers((prev) => [...prev, name])
    );
    socket.on("typing:stop", (name) =>
      setTypingUsers((prev) => prev.filter((u) => u !== name))
    );
    return () => {
      socket.off("typing:start");
      socket.off("typing:stop");
    };
  }, []);

  if (typingUsers.length === 0) return null;

  return (
    <div className="px-4 py-1 text-sm text-gray-500">
      {typingUsers.join(", ")} yazıyor...
    </div>
  );
};