import { useEffect, useState } from 'react';
import { doFetch } from './utils';
import Video from './Video';
function App() {
  const [url, setUrl] = useState();
  const [error, setError] = useState();

  useEffect(() => {
    doFetch('/api/current-id')
      .then((resp) => {
        if (!resp.ok) {
          throw 'Unable to get the latest video. Please refresh and try again';
        }

        return resp.json();
      })
      .then((body) => {
        setUrl(body.url);
      })
      .catch((e) => {
        setError(e);
      });
  }, []);

  if (error) {
    return <h1 className="centered">{error}</h1>;
  }

  if (!url) {
    return <h1 className="centered">Hold tight, getting latest video...</h1>;
  }

  return <Video url={url} />;
}

export default App;
