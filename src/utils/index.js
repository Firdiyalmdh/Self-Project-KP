import { format, parse } from "date-fns";

export const formatDate = (date, outputFormat = 'dd MMMM yyyy') => {
    const inputDate = parse(date, 'dd/MM/yyyy', new Date());
    return format(inputDate, outputFormat);
}

export const getCurrentDate = (outputFormat = "dd/MM/yyyy") => {
    const currentDate = new Date();
    return format(currentDate, outputFormat);
}

export const truncateString = (str, num) => {
    if (str.length <= num) {
        return str
    }
    return str.slice(0, num) + '...'
}