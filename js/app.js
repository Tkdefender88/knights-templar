/* script.js */
document.addEventListener("DOMContentLoaded", () => {
  const jokeContainer = document.getElementById("jokeContainer");
  const jokes = [
    "Why did Pastor Philip bring a ladder to church? Because he wanted to get closer to God!",
    "Philip tried to turn water into Kool-Aid once. It kinda worked.",
    "His sermons are like skateboards — full of sick flips and rad turns.",
    "Pastor Philip's Bible is so well-read, it's got a bookmark for every page!",
    "You know Pastor Philip is serious when he brings out the 'Pastor Philip 3000' — his trusty Bible concordance.",
    "Philip's idea of a wild night is two Bible studies back-to-back.",
    "Why doesn't Pastor Philip use a GPS? Because he always follows God's path!",
    "Pastor Philip once tried to make a TikTok, but ended up with a new youth sermon instead!"
  ];
  const joke = jokes[Math.floor(Math.random() * jokes.length)];
  if (jokeContainer) jokeContainer.innerHTML = `<p><strong>Pastor Joke:</strong> ${joke}</p>`;
  const audio = document.getElementById('philip-audio');
  const playBtn = document.getElementById('play-audio-btn');
  if (playBtn && audio) {
    playBtn.addEventListener('click', () => {
      audio.play().catch(error => {
        console.error('Audio play failed:', error);
      });
    });
  }
});
