import { useState, useEffect } from 'react';

function Video({ url }) {
  const fullUrl = `https://www.facebook.com/plugins/video.php?href=${encodeURIComponent(
    url
  )}&show_text=false&width=auto`;

  const getDimensions = () => {
    const width = document.documentElement.clientWidth;
    const height = document.documentElement.clientHeight;
    var widthScale = width / 1280;
    var heightScale = height / 720;
    var aspectRatioScale = widthScale < heightScale ? widthScale : heightScale;
    return {
      width: 1280 * aspectRatioScale,
      height: 720 * aspectRatioScale,
    };
  };

  const [dim, setDim] = useState(getDimensions());

  useEffect(() => {
    window.addEventListener('resize', onResize);

    return function cleanup() {
      window.removeEventListener('resize', onResize);
    };
  }, []);

  const onResize = () => {
    setDim(getDimensions());
  };

  return (
    <iframe
      title="video"
      width={dim.width}
      height={dim.height}
      src={fullUrl}
      scrolling="no"
      frameBorder="0"
      className="responsive-iframe"
      allow="autoplay; clipboard-write; encrypted-media; picture-in-picture; web-share"
      allowFullScreen
    />
  );
}

export default Video;
