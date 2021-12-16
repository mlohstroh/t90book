export const isLocalhost = () => {
  return window.location.hostname.includes("localhost");
};

export const doFetch = (path, options) => {
  const hostname = isLocalhost() ? "http://localhost:5000" : "";
  return fetch(hostname + path, options);
};
