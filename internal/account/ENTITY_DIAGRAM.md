# ENTITY DIAGRAM AND DEFAULT DATA

```mermaid
erDiagram
    ROLES {
        int id
        string name
        string description
    }
    
    USERS {
        int id
        int role_id
        string name
        string username
        string email
        string phone
        string password
    }
    
    USERS }|--|| ROLES : one_to_many
```


### Current Default Roles Data
1. admin - all access (web admin & desktop client)
2. cashier - order and payment (desktop client)
3. waiter - reservation and order (desktop client)