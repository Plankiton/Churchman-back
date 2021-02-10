# Church App

---



Owner: Pr. Guilherme Henrique 

Number: +55 86 98828 6044

Price: 2500 (5 parcelas)



App para gerenciar a evolução de fiéis, eventos e células em uma igreja.



## Database scheme

```yaml
EventGroup:
   - name
   - description
   -> Events
Event:
   - name
   - description
   -> Covers
   -> Permissions
   -> Parent:Celule|Role
   -> Members
   - event_type:Culto|Default
   - begin
   - end
Cover:
   - image
   - alt_text
Celule:
   - class
   - leader
   - co_leader
   -> Members:Celular
Member:
   - type
   - status
   -> Person
Person:
   - name
   - telefone
   - email
    -> Roles
Role:
   - name
    - description
    -> Permissions
Permission:
    - name
    - description
   - permission
      - create
      - remove
      - update
      - read
    -> ChildTypes
ChildType:
   -> Roles
   -> Celules
   -> Events
   -> Persons

```



## Celule tree



```tree
root                         - pastor
   Celule:root_G12           - 12 persons role:root_G12
      Role:root_G12
         Celule:default
            Persons          - 12 persons role:
```

## Role tree

```tree
Role:root_G12|teacher
   Role:member
      Role:default
```