export default function PricingTR() {
  return (
    <section className="max-w-6xl mx-auto py-20 px-4">
      <h1 className="text-4xl font-heading text-center mb-4">Fiyatlandırma</h1>
      <p className="text-center mb-10">Kurumsal güvenlik, esnek planlar.</p>

      <div className="grid md:grid-cols-3 gap-6">
        {[
          { title: "Başlangıç", price: "₺0", features: ["10 kullanıcı", "Temel destek"] },
          { title: "Profesyonel", price: "₺199/mo", features: ["Sınırsız kullanıcı", "SSO", "Öncelikli destek"] },
          { title: "Kurumsal", price: "İletişim", features: ["SLA", Beyaz etiket, Özel bulut"] },
        ].map((p) => (
          <div key={p.title} className="border rounded-xl p-6">
            <h2 className="text-xl font-semibold mb-2">{p.title}</h2>
            <p className="text-3xl font-bold mb-4">{p.price}</p>
            <ul className="space-y-2">
              {p.features.map((f) => (
                <li key={f} className="flex items-center gap-2">
                  <span className="text-corp-secondary">✓</span>
                  {f}
                </li>
              ))}
            </ul>
            <button className="mt-6 w-full bg-corp-primary text-white py-2 rounded">
              Seç
            </button>
          </div>
        ))}
      </div>
    </section>
  );
}