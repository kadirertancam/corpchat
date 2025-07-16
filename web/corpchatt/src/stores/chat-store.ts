import { create } from "zustand";
import { devtools } from "zustand/middleware";

export interface Message {
  id: string;
  text: string;
  senderId: string;
  senderName: string;
  avatar?: string;
  createdAt: string;
}

interface ChatState {
  messages: Message[];
  addMessage: (msg: Message) => void;
  setMessages: (msgs: Message[]) => void;
}

export const useChatStore = create<ChatState>()(
  devtools(
    (set) => ({
      messages: [],
      addMessage: (msg) =>
        set((state) => ({ messages: [...state.messages, msg] })),
      setMessages: (msgs) => set({ messages: msgs }),
    }),
    { name: "chat-store" }
  )
);