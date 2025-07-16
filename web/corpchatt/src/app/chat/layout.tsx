import { Sidebar } from "@/app/components/sidebar";
import { Header } from "@/app/components/header";

export default function ChatLayout({ children }: { children: React.ReactNode }) {
  return (
    <div className="flex h-screen">
      <Sidebar />
      <main className="flex-1 flex flex-col">
        <Header />
        <section className="flex-1 overflow-y-auto">{children}</section>
      </main>
    </div>
  );
}