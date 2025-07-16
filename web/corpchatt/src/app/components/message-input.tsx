"use client";
import { Send } from "lucide-react";
import { useState } from "react";

export const MessageInput = () => {
  const [text, setText] = useState("");

  const send = async () => {
    if (!text.trim()) return;
    await fetch("/api/messages", {
      method: "POST",
      body: JSON.stringify({ text }),
    });
    setText("");
  };

  return (
    <div className="flex px-4 py-2 border-t">
      <input
        value={text}
        onChange={(e) => setText(e.target.value)}
        onKeyDown={(e) => e.key === "Enter" && send()}
        className="flex-1 px-3 py-2 border rounded-l"
        placeholder="Bir mesaj yazın..."
      />
      <button
        onClick={send}
        className="bg-corp-accent text-white px-4 rounded-r hover:bg-opacity-90"
      >
        <Send size={20} />
      </button>
    </div>
  );
};