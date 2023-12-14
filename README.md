# zssn

## Problem Description

ZSSN (Zombie Survival Social Network). The world as we know it has fallen into an apocalyptic scenario. A laboratory-made virus is transforming human beings and animals into zombies, hungry for fresh flesh.

You, as a zombie resistance member (and the last survivor who knows how to code), was designated to develop a system to share resources between non-infected humans.

## Requirements

You will develop a ***REST API*** (yes, we care about architecture design even in the midst of a zombie apocalypse!), which will store information about the survivors, as well as the resources they own.

In order to accomplish this, the API must fulfill the following use cases:

- **Add survivors to the database**

  A survivor must have a *name*, *age*, *gender* and *last location (latitude, longitude)*.

  A survivor also has an inventory of resources of their own property (which you need to declare when upon the registration of the survivor).

- **Update survivor location**

  A survivor must have the ability to update their last location, storing the new latitude/longitude pair in the base (no need to track locations, just replacing the previous one is enough).

- **Flag survivor as infected**

  In a chaotic situation like that, it's inevitable that a survivor may get contaminated by the virus.  When this happens, we need to flag the survivor as infected.

  An infected survivor cannot trade with others, can't access/manipulate their inventory, nor be listed in the reports (infected people are kinda dead anyway, see the item on reports below).

  **A survivor is marked as infected when at least three other survivors report their contamination.**

  When a survivor is infected, their inventory items become inaccessible (they cannot trade with others).

- **Survivors cannot Add/Remove items from inventory**

  Their belongings must be declared when they are first registered in the system. After that they can only change their inventory by means of trading with other survivors.

  The items allowed in the inventory are described above in the first feature.

- **Trade items**:

  Survivors can trade items among themselves.

  To do that, they must respect the price table below, where the value of an item is described in terms of points.

  Both sides of the trade should offer the same amount of points. For example, 1 Water and 1 Medication (1 x 4 + 1 x 2) is worth 6 ammunition (6 x 1) or 2 Food items (2 x 3).

  The trades themselves need not to be stored, but the items must be transferred from one survivor to the other.

| Item         | Points   |
|--------------|----------|
| 1 Water      | 4 points |
| 1 Food       | 3 points |
| 1 Medication | 2 points |
| 1 Ammunition | 1 point  |

- **Reports**

  The API must offer the following reports:

    1. Percentage of infected survivors.
    1. Percentage of non-infected survivors.
    3. Average amount of each kind of resource by survivor (e.g. 5 waters per survivor)
    4. Points lost because of infected survivor.

---------------------------------------

## My solution

### API

A postman collection can be imported using the `ZSSN.postman_collection.json` file.

- GET `/items`
  - Returns all the available items
- GET `/items/:name`
  - Returns a specific item
- GET `/survivors`
  - Returns all the survivors
- PUT `/survivors`
  - Registers a new survivor
- GET `/survivors/:sid`
  - Returns a specific survivor
- PATCH `/survivors/:sid/status`
  - Reports a survivor status (healthy, infected, dead)
- PATCH `/survivors/:sid/location`
  - Updates a survivor location
- PUT `/survivors/:sid/items/:name`
  - Adds new items to a survivor inventory
  - Query parameters: 
    - quantity: int
- DELETE `/survivors/:sid/items/:name`
  - Removes an existing item to a survivor inventory
  - Query parameters: 
    - quantity: int

Trades endpoint is missing, but all the logic can be seen on the service, which has also unit tests covering the 94.7% of the code.

Report logic is missing



**Original test written by [Akita](https://t.co/W47ODZTOAc)**
**Adapter from [nayra0](https://github.com/nayra0/backend-test/blob/master/README.md)**