# ENTITY DIAGRAM AND DEFAULT DATA

```mermaid
erDiagram
    ROLES {
        int id
        string name
    }
    
    USERS ||--|| ROLES : has_one
    USERS {
        int id
        int role_id
        string name
        string username
        string email
        string phone
        string password
    }
```