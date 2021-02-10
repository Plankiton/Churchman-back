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
   - model_type
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
   -> Church
   -> Parent:Celule|Church
   -> Members:Celular
Member:
   - type
   - status
   -> group:Event|Role|Celule|Church
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
            Persons          - 12 persons
```

### Celule algorithm

Quando a celula tem um G12 como líder, ela começa com o seu gênero e o seu número na lista da sua categoria (`M|F`+`{ord_number}`), celulas descendentes teríam o nome do da célula do seu G12 mais o id da subcelula (`G12_id`+`{ord_number0}..{ord_numberN}`).



## Role tree

```tree
Role:root_G12|teacher
   Role:member
      Role:default
```