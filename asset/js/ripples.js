const img = document.getElementById('firstImg');
const ripples = document
  .getElementsByClassName('ripples')[0]
  .getElementsByTagName('div')[0];

img.addEventListener('mouseenter', () => {
  ripples.classList.add('dot');
});

img.addEventListener('mouseleave', () => {
  ripples.classList.remove('dot');
});
