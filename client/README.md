# NextJS Frontend

Simple NextJS frontend app to host recipe conversions backend. Will have a few different pages to support multiple inputs of various configuration, plus single US / Metric ad hoc calculations.

## Tech Used

- [Typescript](https://www.typescriptlang.org/docs/)
- [Tailwind CSS](https://tailwindcss.com/docs/guides/nextjs)
- Components: [ShadCN](https://ui.shadcn.com/docs)
- Icons: [Hero Icons](https://heroicons.com/)

## TODO

- Fix error handling
  - ~~Single Conv~~
  - Multiple Conv - show toast + stop hanging
- Reset input obj after changing dropdown
  - At least disable until new object populated
- Page Structure
  - Add Top Level routes (recipe converter, measurements, etc.)
  - New Pages
    - Heights/weights
- Refactor
  - Fix card outputs - utilise space better
- API
  - Set up API endpoints
    - Route Handler?
    - React Query?
