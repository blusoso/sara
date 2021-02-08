import React from 'react';
import globalStyles from '@/styles/globalStyles';
import { SWRConfig } from 'swr';
import axios from 'axios';
import NavGlobal from '@/components/Nav/NavGlobal';

function MyApp({ Component, pageProps }) {
  const fetcher = (...arg) =>
    axios(...arg)
      .then((res) => res.data)
      .catch((err) => {
        console.error(err);
        return Promise.reject(error);
      });

  return (
    <React.Fragment>
      <SWRConfig
        value={{
          fetcher,
        }}
      >
        <NavGlobal />
        <Component {...pageProps} />
      </SWRConfig>

      <style jsx global>
        {globalStyles}
      </style>
    </React.Fragment>
  );
}

export default MyApp;
