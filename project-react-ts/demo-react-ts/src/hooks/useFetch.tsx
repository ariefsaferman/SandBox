import { useEffect, useState } from "react";

export default function useFetch<Payload>(url: string): {
    data: Payload | undefined;
    loading: boolean;
    error: boolean;
    errorMessage: Error | undefined
} {
    const [data, setData] = useState<Payload | undefined>(undefined);
    const [loading, setLoading] = useState<boolean>(false); 
    const [error, setError] = useState<boolean>(false); 
    const [errorMessage, setErrorMessage] = useState<Error | undefined>(undefined); 

    useEffect(() => {
        setLoading(true); 
        fetch(url)
         .then(response => {
            if (!response.ok) {
                if (response.status === 404) {
                    throw new Error("Error endpoint")
                }
                console.log(response.statusText); 
                throw new Error(response.statusText);
            }
            return response.json()})
         .then((result: Payload) => {
            setData(result);
        })
         .catch((err) => {
            setError(true);
            setErrorMessage(err);
         })
         .finally(() => {
            setLoading(false); 
         });
    }, [url]);

    return {data, loading, error, errorMessage};
}