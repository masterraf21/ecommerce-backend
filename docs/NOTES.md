## Notes for e Commerce App

- All Requirement are implemented except JWT Authentication. I already created the utility functions for JWT. Will be implemented soon.
- I changed the schema for product, assuming price and total price are interchangeable. Insted to handle multiple producst in a single order i created a composite of product detailing the product called order_detail. Hence i use mongoDB for flexibility in creating data and not bounded by SQL schema restrictions.
