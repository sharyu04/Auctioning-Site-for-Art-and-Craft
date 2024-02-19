# Auctioning-Site-for-Art-and-Craft
Golang Demo Project - Auctioning site for art and craft

# Use Case

1. Users: Seller, Bidder, Admin, Superadmin
2. Sellers and bidders on the website can register using first name, last name, email and password. 
3. Only super admin can register admins 
4. Anyone can log in using their registered email and password.
5. Every user can have the dual role of being both a seller and a bidder. However, a seller cannot bid on their artwork listed for sale.
6. Details about an artwork, such as images, descriptions, starting prices, and the recorded bids, are accessible to all the users.
7. Sellers can list their artwork for auction, with details such as name,description, starting prices, and the duration of the auction. 
8. Sellers have the option to delete their artwork from sale.
9. Bidders can view the artwork on the website and refine their search by applying specific filters to find desired categories.
10. Bidders can place bids on artworks. The bid amount should be greater than the current highest bid. If there are no bids, the bid amount should be greater than the starting price.
11. Bidder can update his bid. The updated amount should be greater than the current highest bid amount.
12. Admin and super admin have all the view accesses and the admin can delete an inappropriate artwork listing.
13. Admin and super admin can view the registered users list with filters.
