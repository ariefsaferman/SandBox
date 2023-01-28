import { useEffect, useState } from "react";

// Using Payload for Generic
export default function useFetch<Payload>(url: string): {
  data: Payload | undefined;
  loading: boolean | undefined;
  error: boolean;
} {
  const [data, setData] = useState<Payload | undefined>(undefined);
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<boolean>(false);

  useEffect(() => {
    setLoading(true);
    fetch(url)
      .then((resp) => resp.json())
      .then((data: Payload) => setData(data))
      .catch((err) => {
        setError(true);
        console.log(err);
      })
      .finally(() => setLoading(false));
  }, [url]);

  return { data, loading, error };
}
