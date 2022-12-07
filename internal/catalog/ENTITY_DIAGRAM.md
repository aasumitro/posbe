# ENTITY DIAGRAM AND DEFAULT DATA

```mermaid
erDiagram
    CATEGORIES {
        int id
        string name
    }
    
    SUBCATEGORIES {
        int id
        int category_id
        string name
    }
    
    UNITS { 
        int id
        string magnitude
        string name
        string symbol
    }
    
    ADDONS {
        int id
        int product_id
        string name
        string description 
        int price
    }
    
    VARIANTS {
        int id
        int product_id
        string name
        string description 
        int price
    }
    
    PRODUCTS {
        int id  
        int unit_id
        int category_id
        int subcategory_id
        string sku
        string image
        json gallery
        string name
        string description
        int price
    }
 
    CATEGORIES ||--|{ SUBCATEGORIES: has_many
    PRODUCTS }|--|| SUBCATEGORIES: has_many
    PRODUCTS ||--|{ ADDONS: one_to_many
    PRODUCTS ||--|{ VARIANTS: one_to_many
    PRODUCTS }|--|| UNITS : one_to_many
    PRODUCTS }|--|| CATEGORIES : one_to_many
```
#### ADDONS:
e.g:
1. EXTRA MILK
2. EXTRA SUGAR
3. EXTRA ...

#### VARIANTS: 
e.g:
1. Tall
2. Grande
3. Venti

#### CATEGORY
1. Beverage
2. Foods

#### SUBCATEGORY:
1. Beverage:
   1. Water
   2. Milk
   3. Tea
   4. Coffee
   5. Sparkling drinks
   6. Juices
   7. Energy drink
   8. Mocktails
   9. Cocktails
   10. Milkshakes
   11. Smoothies
   12. Tonic Water
   13. Beer
   14. Wine
   15. Cider

2. Foods:
   1. Fat
   2. Protein
   3. Dairy
   4. Starchy food
   5. Fruit and vegetables
