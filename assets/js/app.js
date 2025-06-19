/* script.js */
document.addEventListener("DOMContentLoaded", () => {
  const jokeContainer = document.getElementById("jokeContainer");
  const jokes = [
    "Why did Pastor Philip bring a ladder to church? Because he wanted to get closer to God!",
    "Philip tried to turn water into Kool-Aid once. It kinda worked.",
    "His sermons are like skateboards — full of sick flips and rad turns."
  ];
  const joke = jokes[Math.floor(Math.random() * jokes.length)];
  if (jokeContainer) jokeContainer.innerHTML = `<p><strong>Pastor Joke:</strong> ${joke}</p>`;
});
