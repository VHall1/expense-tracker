type Transaction = {
  id: string;
  amount: number;
  description: string;
  notes: string;
  settled_at: Date;
  created_at: Date;
  updated_at: Date;
};

const LEDGER_SERVICE_URL = "http://localhost:8001";
type HTTPMethods = "GET" | "POST";

function makeRequest(path: string, method: HTTPMethods) {
  return fetch(`${LEDGER_SERVICE_URL}${path}`, { method });
}

async function getTransactions(): Promise<Transaction[]> {
  const response = await makeRequest("/transactions", "GET");

  type JSONResponse = {
    data?: Transaction[];
    error?: string;
  };
  const { data, error }: JSONResponse = await response.json();

  if (!response.ok) {
    return Promise.reject(new Error(error ?? "unhandled api error"));
  }

  // return an empty array in the unlikely case the api returns a nullish value
  return data ?? [];
}

export { getTransactions };
