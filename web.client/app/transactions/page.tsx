import { getTransactions } from "@/lib/ledger";

export default async function Transactions() {
  const transactions = await getTransactions();
  console.log(await transactions.json());

  return <p>Hello</p>;
}
