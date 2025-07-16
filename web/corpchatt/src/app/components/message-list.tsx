"use client";

import { useEffect } from "react";
import Image from "next/image";  // buraya al
import { useChatStore } from "@/stores/chat-store";
import { socket } from "@/lib/socket";

export const MessageList = () => {
  const { messages, addMessage, setMessages } = useChatStore();

  useEffect(() => {
    socket.on("message:new", addMessage);
    socket.emit("messages:fetch");
    socket.on("messages:list", setMessages);
    return () => {
      socket.off("message:new", addMessage);
      socket.off("messages:list", setMessages);
    };
  }, [addMessage, setMessages]);

  return (
    <div className="flex-1 p-4 space-y-3 overflow-y-auto">
      {messages.map((msg) => (
        <div key={msg.id} className="flex items-start gap-3">
          <Image
            src={msg.avatar || "/avatar.png"}
            alt={msg.senderName}
            className="rounded-full"
            width={40}
            height={40}
            priority={false}
            style={{ objectFit: "cover" }}
          />
          <div>
            <p className="font-semibold">{msg.senderName}</p>
            <p className="text-sm">{msg.text}</p>
            <p className="text-xs text-gray-400">
              {new Date(msg.createdAt).toLocaleTimeString("tr-TR")}
            </p>
          </div>
        </div>
      ))}
    </div>
  );
};
