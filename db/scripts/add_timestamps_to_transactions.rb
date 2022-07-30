Transaction.all.each do |t|
  t.update!(
    created_at: DateTime.now,
    updated_at: DateTime.now
  )
end
