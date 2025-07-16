"use client";

import { useState } from "react";
import { ChevronLeft, MessageSquare, Settings, Users, Shield } from "lucide-react";
import { useUser } from "@auth0/nextjs-auth0/client";

export const Sidebar = () => {
  const [open, setOpen] = useState(true);
  const { user } = useUser();

  const roles = Array.isArray(user?.["https://corpchat.com/roles"]) 
    ? (user!["https://corpchat.com/roles"] as string[]) 
      : undefined;

      const isAdmin = roles?.includes("admin");

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
          { icon: MessageSquare, label: "Sohbetler", href: "#" },
          { icon: Users, label: "Kullanıcılar", href: "#" },
          { icon: Settings, label: "Ayarlar", href: "#" },
        ].map(({ icon: Icon, label, href }) => (
          <a
            key={label}
            href={href}
            className="flex items-center gap-3 px-2 py-2 rounded hover:bg-white/10"
          >
            <Icon size={20} />
            {open && <span>{label}</span>}
          </a>
        ))}

        {isAdmin && (
          <a
            href="/admin"
            className="flex items-center gap-3 px-2 py-2 rounded hover:bg-white/10"
          >
            <Shield size={20} />
            {open && <span>Yönetim Paneli</span>}
          </a>
        )}
      </nav>
    </aside>
  );
};
