import React from "react";
import { useParams } from "react-router-dom";
import EditForm from "../form/EditForm";
import { IBook } from "../interface/books";
import useFetch from "../hooks/useFetch";

export default function EditBookPage() {
  const params = useParams();
  const { data, callRefresh } = useFetch<IBook>(
    `http://localhost:3004/books/${params.id}`
  );

  return (
    <div className="d-flex vh-100 flex-column align-items-center justify-content-center">
      <EditForm refresh={callRefresh} editedBook={data}></EditForm>
    </div>
  );
}
