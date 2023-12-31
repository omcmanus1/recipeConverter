# NextJS Frontend

Simple NextJS frontend app to host recipe conversions backend. Will have a few different pages to support multiple inputs of various configuration, plus single US / Metric ad hoc calculations.

## Tech Used

- [Typescript](https://www.typescriptlang.org/docs/)
- [Tailwind CSS](https://tailwindcss.com/docs/guides/nextjs)
- Components: [ShadCN](https://ui.shadcn.com/docs)
- Icons: [Hero Icons](https://heroicons.com/)

## TODO

- ~~Add metric > US conversion~~
- Remove conversions from list after adding 
- Page Structure
  - Add Top Level routes (recipe converter, measurements, etc.)
  - New Pages
    - Heights/weights
- Refactor
  - ~~Reusable component for single conversion~~
  - ~~Fix prop drilling~~
  - ~~Refactor volume/weight inputs (reusable object?)~~
  - ~~Abstract inputs component - reuse for single & list conversion~~
  - ~~Make new input list component (multiple items in one card)~~
  - Fix card outputs - utilise space better
  - Add delete button in list component
- API
  - ~~Single fetch request~~
  - Set up API endpoints
    - Route Handler?
    - React Query?
