import { format, parse } from "date-fns";

export const formatDate = (date, outputFormat = 'dd MMMM yyyy') => {
    const inputDate = parse(date, 'dd/MM/yyyy', new Date());
    return format(inputDate, outputFormat);
}