status: check_in, order_placement, print_bill, paid, cancel
order (
    id, cashier_id, shift_id, table_id, room_id,
    date, time_open, time_close, customer,
    brutto, discount, netto, tax, total,
    type, payment, change,
    notes, status, cancel_reason,
    created_at, updated_at
)
order_products (
    id, order_id, product_id, category_id,
    subcategory_id, variant_id,
    name, quantity, price, netto
    notes, created_at, updated_at
)
order_product_addons (
    id, order_id, order_product_id, addon_id
   name, quantity, price, netto,
   notes, created_at, updated_at
)
