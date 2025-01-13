function makeRequest(path: string, method: string) {
  return fetch(`http://localhost:8080${path}`, { method });
}

function getTransactions() {
  return makeRequest("/transactions", "GET");
}

export { getTransactions };
