# ENTITY DIAGRAM AND DEFAULT DATA

```mermaid
erDiagram
    CATEGORIES ||--|{ SUBCATEGORIES: has_many
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
    
    PRODUCTS ||--|{ ADDONS: has_many
    PRODUCTS ||--|{ VARIANTS: has_many
    PRODUCTS ||--|| UNITS : has_one
    PRODUCTS ||--|| CATEGORIES : has_one
    PRODUCTS {
        int id  
        int unit_id
        int category_id
    }
```
#### ADDONS:
1. Beverage
   1. EXTRA MILK
   2. EXTRA SUGAR
   3. EXTRA ...

#### VARIANTS: 
1. Beverage:
   1. Tall
   2. Grande
   3. Venti

2. Foods:
   1. asd
   2. asd
   3. asd

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
 


