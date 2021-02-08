import css from 'styled-jsx/css';

export default css.global`
  html,
  body {
    padding: 0;
    margin: 0;
    font-family: -apple-system, BlinkMacSystemFont, Segoe UI, Roboto, Oxygen,
      Ubuntu, Cantarell, Fira Sans, Droid Sans, Helvetica Neue, sans-serif;
  }

  a {
    color: inherit;
    text-decoration: none;
  }

  * {
    box-sizing: border-box;
  }

  main,
  .main {
    position: relative;
    max-width: 2560px;
    margin: 0 auto;
  }

  .contianer {
    position: relative;
    margin: 0 auto;
    max-width: 1200px;
    padding: 0 22px;
  }
`;
