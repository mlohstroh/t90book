function Video({ url }) {
  const fullUrl = `https://www.facebook.com/plugins/video.php?href=${encodeURIComponent(
    url
  )}&show_text=false&width=auto`;

  return (
    <iframe
      title="test"
      width="auto"
      height="100%"
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
