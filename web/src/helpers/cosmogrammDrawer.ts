//создает сегменты кольца
function generateRingSectors(outerRadius: number, innerRadius: number, numSectors: number, cx: number, cy: number, hoverScale: number, textList: number[]) {
  const sectorAngle = 360 / numSectors;
  const sectors = [];
  const offsetAngle = 0;

  for (let i = 0; i < numSectors; i++) {
    const startAngle = -((i * sectorAngle + offsetAngle) * Math.PI) / 180;
    const endAngle = -(((i + 1) * sectorAngle + offsetAngle) * Math.PI) / 180;

    const pathData = `
      M ${cx + outerRadius * Math.cos(startAngle)} ${cy + outerRadius * Math.sin(startAngle)}
      A ${outerRadius} ${outerRadius} 0 0 0 ${cx + outerRadius * Math.cos(endAngle)} ${cy + outerRadius * Math.sin(endAngle)}
      L ${cx + innerRadius * Math.cos(endAngle)} ${cy + innerRadius * Math.sin(endAngle)}
      A ${innerRadius} ${innerRadius} 0 0 1 ${cx + innerRadius * Math.cos(startAngle)} ${cy + innerRadius * Math.sin(startAngle)}
      Z
    `;

    const pathDataHover = `
      M ${cx + (outerRadius + hoverScale) * Math.cos(startAngle)} ${cy + (outerRadius + hoverScale) * Math.sin(startAngle)}
      A ${outerRadius + hoverScale} ${outerRadius + hoverScale} 0 0 0 ${cx + (outerRadius + hoverScale) * Math.cos(endAngle)} ${cy + (outerRadius + hoverScale) * Math.sin(endAngle)}
      L ${cx + (innerRadius - hoverScale) * Math.cos(endAngle)} ${cy + (innerRadius - hoverScale) * Math.sin(endAngle)}
      A ${innerRadius - hoverScale} ${innerRadius - hoverScale} 0 0 1 ${cx + (innerRadius - hoverScale) * Math.cos(startAngle)} ${cy + (innerRadius - hoverScale) * Math.sin(startAngle)}
      Z
    `;

    const path = document.createElementNS("http://www.w3.org/2000/svg", "path");

    path.setAttribute("d", pathData);
    path.setAttribute('fill', '#780986');
    path.classList.add("piece");


    const text = document.createElementNS("http://www.w3.org/2000/svg", "text");

    text.textContent = String.fromCharCode(textList[i]);
    text.setAttribute("x", (cx + (outerRadius - (outerRadius - innerRadius) / 2) * Math.cos((startAngle + endAngle) / 2)).toString());
    text.setAttribute("y", (cy + (outerRadius  - (outerRadius - innerRadius) / 2) * Math.sin((startAngle + endAngle) / 2)).toString());
    text.setAttribute('font-size', '40px');
    text.classList.add('piece-text');


    const pathHover = () => {
      path.setAttribute("d", pathDataHover);
      path.setAttribute('fill', '#bb13d1');
      text.setAttribute('font-size', '45px');
    };

    const pathUnhover = () => {
      path.setAttribute("d", pathData);
      path.setAttribute('fill', '#780986');
      text.setAttribute('font-size', '40px');
    }
    
    path.addEventListener("mouseover", pathHover);
    text.addEventListener("mouseover", pathHover);
    
    path.addEventListener("mouseout", pathUnhover);
    text.addEventListener("mouseout", pathUnhover);


    sectors.push(path);
    sectors.push(text);
  }

  return sectors;
}

//создает круг
function createCircle(centerX: number, centerY: number, radius: number, classlist: string[]) {
  const circle = document.createElementNS("http://www.w3.org/2000/svg", "circle");

  circle.setAttribute("cx", centerX.toString());
  circle.setAttribute("cy", centerY.toString());
  circle.setAttribute("r", radius.toString());
  classlist.forEach((item) => circle.classList.add(item));

  return circle;
}

//создает прерывистые линии на кругу, паттерн: [линия, пробел, линия, пробел, ...]
function generateDashArray(radius: number, pattern: number[]){
  const circumference = 2 * Math.PI * radius;
  const sum = pattern.reduce((sum: number, number: number) => sum + number, 0);
  const segmentLengths = pattern.map(length => (length / sum) * circumference);
  
  // Создаем строку stroke-dasharray
  let dashArray = '';
  segmentLengths.forEach((length, index) => {
    dashArray += length;
    if (index < segmentLengths.length - 1) {
      dashArray += ' ';
    }
  });

  return dashArray;
}

export function generateCosmogramm(container: HTMLElement, width: number): number{
  const radius = Math.floor(width / 2);
  const step = 12;
  container.appendChild(createCircle(radius, radius, radius - 4 - step * 3, ["circle", "circle-big"]));
  
  generateRingSectors(radius - 4 - step * 5, radius - 4 - step * 11, 12, radius, radius, 7, [97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108]).forEach((item) => container.appendChild(item));
  
  container.appendChild(createCircle(radius, radius, radius - 4 - step * 13, ["circle", "circle-small"]));
  
  
  const circle1 = createCircle(radius, radius, radius - 4 - step * 2, ['circle-spin', 'circle-spin-1']);
  circle1.setAttribute('stroke-dasharray', generateDashArray(radius - 3 - step * 2, [1, 5, 3, 2, 5, 4, 3, 1]));
  container.appendChild(circle1);
  
  const circle2 = createCircle(radius, radius, radius - 4 - step, ['circle-spin', 'circle-spin-2']);
  circle2.setAttribute('stroke-dasharray', generateDashArray(radius - 4 - step, [5, 3, 3, 2, 1, 1, 4, 4, 2, 5]));
  container.appendChild(circle2);
  
  const circle3 = createCircle(radius, radius, radius - 4, ['circle-spin', 'circle-spin-3']);
  circle3.setAttribute('stroke-dasharray', generateDashArray(radius - 4, [1, 2, 1, 2, 2, 5, 4, 3, 5, 5, 4, 1, 4, 3]));
  container.appendChild(circle3);  

  return  radius - 4 - step * 13; // small circle radius
}

export function drawLine(container:HTMLElement, point1: Point, point2: Point,  classes: string[]){
  const line = document.createElementNS("http://www.w3.org/2000/svg", "line");

  line.setAttribute("x1", point1.x.toString());
  line.setAttribute("y1", point1.y.toString());
  line.setAttribute("x2", point2.x.toString());
  line.setAttribute("y2", point2.y.toString());

  classes.forEach(item => line.classList.add(item));

  container.appendChild(line);
}

export interface Point {
  x: number;
  y: number;
}

export function calculatePointCoordinates(xc: number, yc: number, radius: number, angle: number, offsetAngle: number = 0): Point {
  const radAngle = (angle + offsetAngle) * (Math.PI / 180);
  return {
    x: xc + radius * Math.cos(radAngle),
    y: yc - radius * Math.sin(radAngle),
  } as Point;
}

export interface ILineType {
  show: boolean,
  class: string,
}

export function checkAngleDifference(angle1: number, angle2: number, accuracy: number = 10):ILineType {
  const diff = Math.abs(angle1 - angle2);
  if(360 - accuracy <= diff || diff <= accuracy) return {show: true, class: 'line-dark'}
  if(60 - accuracy <= diff && diff <= 60 + accuracy) return {show: true, class: 'line-good'}
  if(90 - accuracy <= diff && diff <= 90 + accuracy) return {show: true, class: 'line-bad'}
  if(120 - accuracy <= diff && diff <= 120 + accuracy) return {show: true, class: 'line-good'}
  if(180 - accuracy <= diff && diff <= 180 + accuracy) return {show: true, class: 'line-bad'}
  if(240 - accuracy <= diff && diff <= 240 + accuracy) return {show: true, class: 'line-good'}
  if(270 - accuracy <= diff && diff <= 270 + accuracy) return {show: true, class: 'line-bad'}
  if(300 - accuracy <= diff && diff <= 300 + accuracy) return {show: true, class: 'line-good'}
  return {show: false, class: ''};
}