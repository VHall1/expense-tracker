import { getTransactions, Transaction } from "@/lib/ledger";

export default async function Transactions() {
  const transactionsResponse = await getTransactions();
  const transactions: Transaction[] = await transactionsResponse.json();

  return transactions.map((transaction) => (
    <div key={transaction.id}>{transaction.description}</div>
  ));
}
