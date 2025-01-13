import { getTransactions } from "@/lib/ledger-service";

export default async function Transactions() {
  const transactions = await getTransactions();

  return transactions.map((transaction) => (
    <div key={transaction.id}>{transaction.description}</div>
  ));
}
