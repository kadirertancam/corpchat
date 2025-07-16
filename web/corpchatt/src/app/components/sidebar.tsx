"use client";
import { useState } from "react";
import { ChevronLeft, MessageSquare, Settings, Users } from "lucide-react";

export const Sidebar = () => {
  const [open, setOpen] = useState(true);

  return (
    <aside
      className={`transition-all duration-300 bg-corp-primary text-white flex flex-col ${
        open ? "w-64" : "w-20"
      }`}
    >
      <div className="flex items-center justify-between p-4">
        <span className={`font-heading text-xl ${!open && "hidden"}`}>CorpChat</span>
        <button onClick={() => setOpen(!open)}>
          <ChevronLeft className={`${!open && "rotate-180"}`} />
        </button>
      </div>

      <nav className="flex-1 px-2 space-y-2">
        {[
          { icon: MessageSquare, label: "Sohbetler" },
          { icon: Users, label: "Kullanıcılar" },
          { icon: Settings, label: "Ayarlar" },
        ].map(({ icon: Icon, label }) => (
          <a
            key={label}
            href="#"
            className="flex items-center gap-3 px-2 py-2 rounded hover:bg-white/10"
          >
            <Icon size={20} />
            {open && <span>{label}</span>}
          </a>
        ))}
      </nav>
    </aside>
  );
};