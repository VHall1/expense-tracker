export interface Transaction {
  id: string;
  amount: number;
  description: string;
  notes: string;
  settled_at: Date;
  created_at: Date;
  updated_at: Date;
}

function makeRequest(path: string, method: string) {
  return fetch(`http://localhost:8080${path}`, { method });
}

function getTransactions() {
  return makeRequest("/transactions", "GET");
}

export { getTransactions };
