import HeroAnimation from "@/components/AnimatedHero";
import NewScene from "@/components/NewScene";
import Scene from "@/components/Scene";

export default function Home() {
  return (
    <section className="w-full mx-auto  py-12">
      {/* <NewScene /> */}
      <Scene />
      {/* <div className="text-center mb-12">
        <h1 className="text-4xl font-extrabold text-(--foreground)">
          Umów się z <span className="text-(--primary)">Sałatą</span>
        </h1>
        <p className="mt-4 text-lg text-(--foreground) opacity-70">
          Apka do zmawiania się ze znajomymi.
        </p>
      </div>

      <HeroAnimation />

      <div className="h-screen flex items-center justify-center bg-(--background)">
        <p className="text-(--foreground)">coś tu bedzie</p>
      </div> */}
    </section>
  );
}
