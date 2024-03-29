import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

type SelectProps = {
  handleChange?: (inp: string, val?: string) => void;
  placeholder?: string;
  selectContent: string[];
  disabled?: boolean;
  classNames?: string;
};

export default function SelectSh({
  handleChange,
  placeholder,
  selectContent,
  disabled = false,
  classNames,
}: SelectProps) {
  return (
    <div className={`mb-1 ${classNames}`}>
      <Select onValueChange={handleChange}>
        <SelectTrigger disabled={disabled}>
          <SelectValue placeholder={placeholder} />
        </SelectTrigger>
        <SelectContent>
          {selectContent.map((item) => {
            return (
              <SelectItem key={`${item}_key`} value={item}>
                {item.charAt(0).toUpperCase() + item.slice(1)}
              </SelectItem>
            );
          })}
        </SelectContent>
      </Select>
    </div>
  );
}
