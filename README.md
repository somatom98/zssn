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

## Notes

1. Please use one of the following languages: Go or Python. You are free to use whatever framework you'd like.
2. No authentication is needed (it's a zombie apocalypse, no one will try to hack a system while running from a horde of zombies);
3. We still care about proper programming and architecture techniques, you must showcase that you're worthy of surving the zombie apocalypse through the sheer strength of your skills; 
4. You must write at least some automated tests;
4. Aimple is better. Do not over-engineer or over-architect your solution. We don't need to see the *best* or the *most complete* solution, but rather a production-ready solution that you would be able to deploy to a cloud provider;
4. Don't forget to make at least a minimal documentation of the API endpoints and how to use them;
5. From the problem description above you can either do a very bare bones solution or add optional features that are not described. Use your time wisely; the absolute optimal solution might take too long to be effective in the apocalypse, so you must come up with the best possible solution that will hold up within the least ammount of time and still be able to showcase your skills in order to prove your worth.
7. **Optional: Deploy your solution to Cloud Run and connect with Cloud SQL db (you get 300$ credit when you signup).**

## Q&A

> Where should I send back the result when I'm done?

Create a public git repository on Gitlab/Github and send us an email with the link to j.acosta@neuronsinc.com. Please make sure to create a proper Readme file with the link to the Cloud Run instance so we can test your solution out.

> What if I have a question?

You can always reach out to j.acosta@neuronsinc.com if you have questions or concerns about this test.

**Original test written by [Akita](https://t.co/W47ODZTOAc)**
**Adapter from [nayra0](https://github.com/nayra0/backend-test/blob/master/README.md)**