class CardService {
  createNewsCard(news) {
    const card = document.createElement("div");

    card.className = "news-card";

    if (picture) {
      const img = document.createElement("img");
      img.src = news.picture;
      img.alt = news.title;
      card.appendChild(img);
    }

    const author = document.createElement('h5');
    author.textContent = news.author;
    card.appendChild(author);

    const createdAt = document.createElement('h5');
    createdAt.textContent = news.createdAt;
    card.appendChild(createdAt);

    const h2 = document.createElement('h2');
    h2.textContent = news.title;
    card.appendChild(h2);

    return card;
  }
}

module.exports = new CardService();
