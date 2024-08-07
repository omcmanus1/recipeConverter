import Navigation from "@/components/navigation";

export default function Layout({ children }: { children: React.ReactNode }) {
  const recipeTabs = [
    { title: "Convert US To Metric", path: "/recipes/usa" },
    { title: "Convert Metric To US", path: "/recipes/metric" },
    { title: "Convert List", path: "/recipes/list" },
  ];

  return (
    <div className="flex flex-col items-center text-center">
      <Navigation tabs={recipeTabs}></Navigation>
      <div>{children}</div>
    </div>
  );
}
