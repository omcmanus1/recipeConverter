import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip";
import { TooltipArrow } from "@radix-ui/react-tooltip";

export default function TooltipSh({
  title,
  tooltip,
}: {
  title: string;
  tooltip: string;
}) {
  return (
    <TooltipProvider>
      <Tooltip>
        <TooltipTrigger className="cursor-not-allowed">{title}</TooltipTrigger>
        <TooltipContent sideOffset={10}>
          <p>{tooltip}</p>
          <TooltipArrow />
        </TooltipContent>
      </Tooltip>
    </TooltipProvider>
  );
}
