/*!
* Start Bootstrap - Personal v1.0.1 (https://startbootstrap.com/template-overviews/personal)
* Copyright 2013-2023 Start Bootstrap
* Licensed under MIT (https://github.com/StartBootstrap/startbootstrap-personal/blob/master/LICENSE)
*/
// This file is intentionally blank
// Use this file to add JavaScript to your project

document.getElementById('contactForm').addEventListener('submit', function(event) {
    event.preventDefault(); // 폼의 기본 제출 동작을 막음

    // 입력 필드 값 가져오기
    const name = document.getElementById('name').value;
    const email = document.getElementById('email').value;
    const phone = document.getElementById('phone').value;
    const message = document.getElementById('message').value;

    // 폼 데이터 객체 생성
    const data = {
        name: name,
        email: email,
        phone: phone,
        message: message
    };

    // 서버로 POST 요청 보내기
    fetch('/submit', { // 여기에 실제 서버 URL 입력
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
    .then(response => response.json())
    .then(data => {
        console.log('Success:', data);
        // 성공 시 처리 로직 추가
        alert("Form submitted successfully!");
    })
    .catch((error) => {
        console.error('Error:', error);
        // 오류 시 처리 로직 추가
        alert("An error occurred. Please try again.");
    });
});


/*
document.addEventListener("DOMContentLoaded", function () {
    // API로부터 JSON 데이터 가져오기
    fetch('/resume')
        .then(response => response.json())
        .then(data => {
            // 데이터를 HTML 요소에 삽입하기
            const experience1 = data.experience[0];
            document.getElementById('experience1-period').textContent = experience1.period;
            document.getElementById('experience1-role').textContent = experience1.role;
            document.getElementById('experience1-company').textContent = experience1.company;
            document.getElementById('experience1-location').textContent = experience1.location;
            document.getElementById('experience1-description').textContent = experience1.description;

            // 두 번째 경험 (또는 더 많은 경험을 추가하려면 유사한 코드 작성)
            const experience2 = data.experience[1];
            document.getElementById('experience2-period').textContent = experience2.period;
            document.getElementById('experience2-role').textContent = experience2.role;
            document.getElementById('experience2-company').textContent = experience2.company;
            document.getElementById('experience2-location').textContent = experience2.location;
            document.getElementById('experience2-description').textContent = experience2.description;

            const skills = data.skills; // 스킬 배열 데이터 접근
            const container = document.getElementById('skillsContainer'); // 컨테이너 요소 가져오기

            const languages = data.languages; // 언어스킬 배열에 접근
            const containerlang = document.getElementById('languagesContainer')

            // 각 스킬 항목을 처리
            skills.forEach(skill => {
                // 각 스킬을 담을 새로운 div 요소 생성
                const skillDiv = document.createElement('div');
                skillDiv.className = 'col mb-4 mb-md-0'; // Bootstrap 클래스 추가
                skillDiv.innerHTML = `
                    <div class="d-flex align-items-center bg-light rounded-4 p-3 h-100">
                        ${skill}
                    </div>
                `;
                container.appendChild(skillDiv); // 컨테이너에 추가
            });

            languages.forEach(language => {
                // 각 스킬을 담을 새로운 div 요소 생성
                const languageDiv = document.createElement('div');
                languageDiv.className = 'col mb-4 mb-md-0'; // Bootstrap 클래스 추가
                languageDiv.innerHTML = `
                    <div class="d-flex align-items-center bg-light rounded-4 p-3 h-100">
                        ${language}
                    </div>
                `;
                containerlang.appendChild(languageDiv); // 컨테이너에 추가
            });


        })
        .catch(error => console.error('Error fetching data:', error));
});
*/
document.addEventListener("DOMContentLoaded", function () {
    let allExperienceData = [];
    let currentIndex = 0;
    const experienceContainer = document.getElementById('experienceContainer');
    const loadMoreButton = document.getElementById('loadMoreButton');
    const experiencesPerPage = 2; // 한 번에 표시할 경험의 수

    // API로부터 JSON 데이터 가져오기
    fetch('/returnresume')
        .then(response => response.json())
        .then(data => {

            allExperienceData = data.experience;
            displayExperience(); // 처음 페이지 로드 시 첫 두 개의 경험 표시

            const skills = data.skills; // 스킬 배열 데이터 접근
            const container = document.getElementById('skillsContainer'); // 컨테이너 요소 가져오기

            const languages = data.languages; // 언어스킬 배열에 접근
            const containerlang = document.getElementById('languagesContainer')

            // 각 스킬 항목을 처리
            skills.forEach(skill => {
                // 각 스킬을 담을 새로운 div 요소 생성
                const skillDiv = document.createElement('div');
                skillDiv.className = 'col mb-4 mb-md-0'; // Bootstrap 클래스 추가
                skillDiv.innerHTML = `
                    <div class="d-flex align-items-center bg-light rounded-4 p-3 h-100">
                        ${skill}
                    </div>
                `;
                container.appendChild(skillDiv); // 컨테이너에 추가
            });

            languages.forEach(language => {
                // 각 스킬을 담을 새로운 div 요소 생성
                const languageDiv = document.createElement('div');
                languageDiv.className = 'col mb-4 mb-md-0'; // Bootstrap 클래스 추가
                languageDiv.innerHTML = `
                    <div class="d-flex align-items-center bg-light rounded-4 p-3 h-100">
                        ${language}
                    </div>
                `;
                containerlang.appendChild(languageDiv); // 컨테이너에 추가
            });

        })
        .catch(error => console.error('Error fetching data:', error));

    function displayExperience() {
        const endIndex = Math.min(currentIndex + experiencesPerPage, allExperienceData.length);
        
        for (let i = currentIndex; i < endIndex; i++) {
            const experience = allExperienceData[i];
            
            const experienceCard = ` 
                <div class="card shadow border-0 rounded-4 mb-5">
                    <div class="card-body p-5">
                        <div class="row align-items-center gx-5">
                            <div class="col text-center text-lg-start mb-4 mb-lg-0">
                                <div class="bg-light p-4 rounded-4">
                                    <div class="text-primary fw-bolder mb-2">${experience.period}</div>
                                    <div class="small fw-bolder">${experience.role}</div>
                                    <div class="small text-muted">${experience.company}</div>
                                    <div class="small text-muted">${experience.location}</div>
                                </div>
                            </div>
                            <div class="col-lg-8">
                                <div>${experience.description}</div>
                            </div>
                        </div>
                    </div>
                </div>
            `;
            
            experienceContainer.insertAdjacentHTML('beforeend', experienceCard);
        }
        
        currentIndex = endIndex;
        
        // 모든 경험을 다 표시했으면 "더보기" 버튼 숨기기
        if (currentIndex >= allExperienceData.length) {
            loadMoreButton.style.display = 'none';
        }
    }

    // "더보기" 버튼 클릭 이벤트
    loadMoreButton.addEventListener('click', function () {
        displayExperience();
    });
});
