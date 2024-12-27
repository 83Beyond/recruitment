package ui

import (
	"fyne.io/fyne/v2"
)

func getSearchIcon() fyne.Resource {
	var search = &fyne.StaticResource{
		StaticName:    "Search",
		StaticContent: []byte("<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" t=\"1735038445685\" class=\"icon\" viewBox=\"0 0 1024 1024\" version=\"1.1\" p-id=\"12472\" width=\"200\" height=\"200\"><path d=\"M469.333333 768c-166.4 0-298.666667-132.266667-298.666666-298.666667s132.266667-298.666667 298.666666-298.666666 298.666667 132.266667 298.666667 298.666666-132.266667 298.666667-298.666667 298.666667z m0-85.333333c119.466667 0 213.333333-93.866667 213.333334-213.333334s-93.866667-213.333333-213.333334-213.333333-213.333333 93.866667-213.333333 213.333333 93.866667 213.333333 213.333333 213.333334z m251.733334 0l119.466666 119.466666-59.733333 59.733334-119.466667-119.466667 59.733334-59.733333z\" fill=\"#ffffff\" p-id=\"12473\"/></svg>"),
	}
	return search
}

func getBankIcon() fyne.Resource {
	var finance = &fyne.StaticResource{
		StaticName:    "Bank",
		StaticContent: []byte("<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" t=\"1735149875774\" class=\"icon\" viewBox=\"0 0 1024 1024\" version=\"1.1\" p-id=\"21977\" width=\"200\" height=\"200\"><path d=\"M512 1024c-56.096-89.344-97.184-137.728-123.296-145.184C201.216 825.248 64 652.672 64 448 64 200.576 264.576 0 512 0s448 200.576 448 448c0 203.168-135.232 374.72-320.608 429.632-26.88 7.968-69.376 56.768-127.392 146.368z m44-586.688h85.888c14.4 0 26.112-11.584 26.112-25.92a26.016 26.016 0 0 0-26.112-25.92h-70.144l37.28-64.064a25.824 25.824 0 0 0-9.568-35.424 26.24 26.24 0 0 0-35.712 9.504l-35.84 61.568-35.84-61.568a26.24 26.24 0 0 0-35.68-9.504 25.824 25.824 0 0 0-9.6 35.424l37.344 64.064H414.08a26.016 26.016 0 0 0-26.112 25.92c0 14.336 11.712 25.92 26.112 25.92h85.888v25.952h-85.888a26.016 26.016 0 0 0-26.112 25.92c0 14.304 11.712 25.92 26.112 25.92h85.888v49.984c0 15.36 12.544 27.808 28 27.808a27.904 27.904 0 0 0 28-27.808v-49.984h85.888c14.4 0 26.112-11.616 26.112-25.92a26.016 26.016 0 0 0-26.112-25.92h-85.888v-25.92z m224-147.2v-0.064h20.48a51.072 51.072 0 0 0 48.224-34.144 50.592 50.592 0 0 0-31.648-64.448L546.24 99.008a56.416 56.416 0 0 0-36.448 0l-270.848 92.48c-20.576 7.04-34.4 26.24-34.432 47.84a50.88 50.88 0 0 0 51.04 50.752h20.48v302.816H248c-30.912 0-56 24.864-56 55.552S216.864 704 247.552 704h560.896c30.688 0 55.552-24.864 55.552-55.552 0-30.72-25.088-55.552-56-55.552h-28V290.08zM528 620.64c-108.224 0-195.968-87.04-196-194.464 0-107.392 87.744-194.464 196-194.464s196.032 87.072 196 194.464c0 107.392-87.744 194.464-196 194.464z\" fill=\"#8a8a8a\" p-id=\"21978\"/></svg>"),
	}
	return finance
}

func getInsuranceIcon() fyne.Resource {
	var insurance = &fyne.StaticResource{
		StaticName:    "Insurance",
		StaticContent: []byte("<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" t=\"1735149996726\" class=\"icon\" viewBox=\"0 0 1024 1024\" version=\"1.1\" p-id=\"23230\" width=\"200\" height=\"200\"><path d=\"M441.6 306.8L403 288.6c-3.3-1.6-7.3 0.2-8.4 3.7-17.5 58.5-45.2 110.1-82.2 153.6-1.3 1.6-1.8 3.7-1.2 5.6l13.2 43.5c1.3 4.4 7 5.7 10.2 2.4 7.7-8.1 15.4-16.9 23.1-26V656c0 4.4 3.6 8 8 8H403c4.4 0 8-3.6 8-8V393.1c13.8-25.1 25.1-51.7 33.6-79 1-2.9-0.3-6-3-7.3z m26.8 9.2v127.2c0 4.4 3.6 8 8 8h65.9v18.6h-94.9c-4.4 0-8 3.6-8 8v35.6c0 4.4 3.6 8 8 8h55.1c-19.1 30.8-42.4 55.7-71 76-2.6 1.8-3.3 5.4-1.6 8.1l22.8 36.5c1.9 3.1 6.2 3.8 8.9 1.4 31.6-26.8 58.7-62.9 80.6-107.6v120c0 4.4 3.6 8 8 8h36.2c4.4 0 8-3.6 8-8V536c21.3 41.7 47.5 77.5 78.1 106.9 2.6 2.5 6.8 2.1 8.9-0.7l26.3-35.3c2-2.7 1.4-6.5-1.2-8.4-30.5-22.6-54.2-47.8-72.3-76.9h59c4.4 0 8-3.6 8-8V478c0-4.4-3.6-8-8-8h-98.8v-18.6h66.7c4.4 0 8-3.6 8-8V316c0-4.4-3.6-8-8-8H476.4c-4.4 0-8 3.6-8 8z m51.5 42.8h97.9v41.6h-97.9v-41.6z\" p-id=\"23231\" fill=\"#8a8a8a\"/><path d=\"M866.9 169.9L527.1 54.1C523 52.7 517.5 52 512 52s-11 0.7-15.1 2.1L157.1 169.9c-8.3 2.8-15.1 12.4-15.1 21.2v482.4c0 8.8 5.7 20.4 12.6 25.9L499.3 968c3.5 2.7 8 4.1 12.6 4.1s9.2-1.4 12.6-4.1l344.7-268.6c6.9-5.4 12.6-17 12.6-25.9V191.1c0.2-8.8-6.6-18.3-14.9-21.2zM810 654.3L512 886.5 214 654.3V226.7l298-101.6 298 101.6v427.6z\" p-id=\"23232\" fill=\"#8a8a8a\"/></svg>"),
	}
	return insurance
}

func getBigTechIcon() fyne.Resource {
	var bigTech = &fyne.StaticResource{
		StaticName:    "BigTech",
		StaticContent: []byte("<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" t=\"1735043685860\" class=\"icon\" viewBox=\"0 0 1024 1024\" version=\"1.1\" p-id=\"15798\" width=\"200\" height=\"200\"><path d=\"M878.514 154.502a514.944 514.944 0 0 0-44.367-40.458C746.166 42.734 634.074 0.006 512 0.006c-122.074 0-234.165 42.728-322.146 114.038a515.27 515.27 0 0 0-44.367 40.458C55.471 246.774 0.006 372.91 0.006 512s55.464 265.225 145.48 357.497a515.133 515.133 0 0 0 44.367 40.459c87.981 71.309 200.073 114.037 322.146 114.037 122.074 0 234.166-42.729 322.146-114.038a514.81 514.81 0 0 0 44.367-40.458c90.016-92.272 145.48-218.407 145.48-357.497 0.002-139.09-55.463-265.226-145.478-357.498zM333.892 96.443c-20.74 25.117-39.522 54.504-55.894 87.397a448.69 448.69 0 0 1-42.653-29.339c30.209-23.444 63.186-42.883 98.547-58.058z m-144.038 98.518a513.169 513.169 0 0 0 64.095 44.336c-26.688 71.086-43.247 153.969-46.285 242.938H60.967c3.257-50.527 14.83-99.566 34.534-146.153 22.296-52.712 54.034-100.17 94.353-141.121z m-0.001 634.077c-40.318-40.951-72.057-88.409-94.352-141.12-19.704-46.588-31.276-95.628-34.534-146.155h146.697c3.037 88.969 19.597 171.852 46.285 242.938a513.187 513.187 0 0 0-64.096 44.337z m45.491 40.459a449.003 449.003 0 0 1 42.653-29.339c16.371 32.893 35.154 62.28 55.894 87.397-35.36-15.174-68.337-34.613-98.547-58.058z m246.892 90.899c-18.331-4.441-36.453-12.98-54.15-25.552-29.035-20.625-56.73-52.268-80.094-91.506-5.868-9.855-11.435-20.13-16.716-30.775a454.26 454.26 0 0 1 4.807-2.067c46.587-19.704 95.626-31.276 146.153-34.534v184.434z m0-244.535c-61.714 3.539-120.501 18-174.452 41.492-6.991-19.318-13.17-39.578-18.494-60.716-12.408-49.268-19.628-101.244-21.58-154.875h214.525v174.099z m0-233.626H267.711c1.951-53.63 9.171-105.606 21.579-154.874 5.323-21.138 11.502-41.398 18.494-60.718 53.951 23.493 112.738 37.954 174.452 41.493v174.099z m0-234.198c-50.527-3.257-99.566-14.83-146.153-34.534-1.607-0.68-3.209-1.37-4.807-2.067 5.281-10.645 10.848-20.919 16.716-30.774 23.364-39.239 51.06-70.881 80.094-91.506 17.697-12.572 35.819-21.11 54.15-25.551v184.432z m446.262 88.045c19.705 46.587 31.277 95.626 34.535 146.153H816.336c-3.038-88.969-19.598-171.852-46.285-242.938a513.187 513.187 0 0 0 64.096-44.337c40.318 40.952 72.056 88.41 94.351 141.122z m-139.842-181.58a449.145 449.145 0 0 1-42.654 29.339c-16.371-32.893-35.154-62.281-55.895-87.398 35.362 15.174 68.338 34.613 98.549 58.059zM541.764 63.604c18.33 4.44 36.453 12.979 54.15 25.552 29.035 20.625 56.73 52.268 80.094 91.506 5.869 9.855 11.436 20.13 16.717 30.775a445.44 445.44 0 0 1-4.807 2.066c-46.588 19.705-95.627 31.277-146.154 34.534V63.604z m0 244.533c61.713-3.539 120.501-18 174.453-41.492 6.99 19.319 13.17 39.579 18.494 60.717 12.406 49.268 19.627 101.243 21.578 154.874H541.764V308.137z m0 233.626h214.525c-1.951 53.631-9.172 105.607-21.578 154.875-5.324 21.138-11.504 41.396-18.494 60.716-53.952-23.492-112.74-37.953-174.453-41.492V541.763z m54.15 393.081c-17.697 12.572-35.82 21.111-54.15 25.552V775.961c50.527 3.257 99.566 14.829 146.154 34.534a475.96 475.96 0 0 1 4.807 2.066c-5.281 10.646-10.848 20.921-16.717 30.776-23.363 39.239-51.059 70.882-80.094 91.507z m94.193-7.287c20.74-25.118 39.523-54.506 55.895-87.399a448.577 448.577 0 0 1 42.654 29.339c-30.211 23.445-63.187 42.885-98.549 58.06z m238.391-239.639c-22.295 52.711-54.033 100.169-94.352 141.12a513.113 513.113 0 0 0-64.096-44.337c26.688-71.087 43.248-153.97 46.285-242.938h146.697c-3.257 50.527-14.829 99.567-34.534 146.155z\" fill=\"#8a8a8a\" p-id=\"15799\"/></svg>"),
	}
	return bigTech
}
