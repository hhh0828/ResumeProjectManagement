document.addEventListener("DOMContentLoaded", function () {
    
    // 서버에서 스킬 및 언어 데이터를 가져옴
    fetch('/returnskillang')
        .then(response => response.json())
        .then(data => {
            // 각 스킬 항목을 처리
            const skills = data.skills; // 스킬 배열 데이터 접근
            const container = document.getElementById('skillsContainer'); // 컨테이너 요소 가져오기
            const languages = data.languages; // 언어스킬 배열에 접근
            const containerlang = document.getElementById('languagesContainer');   

            // 스킬 목록 렌더링
            if (container) {
                skills.forEach(skill => {
                    const skillDiv = document.createElement('div');
                    skillDiv.className = 'col mb-4 mb-md-0';
                    skillDiv.innerHTML = `
                        <div class="d-flex align-items-center bg-light rounded-4 p-3 h-100">
                            ${skill}
                        </div>
                    `;
                    container.appendChild(skillDiv);
                });
            }

            // 언어 목록 렌더링
            if (containerlang) {
                languages.forEach(language => {
                    const languageDiv = document.createElement('div');
                    languageDiv.className = 'col mb-4 mb-md-0';
                    languageDiv.innerHTML = `
                        <div class="d-flex align-items-center bg-light rounded-4 p-3 h-100">
                            ${language}
                        </div>
                    `;
                    containerlang.appendChild(languageDiv);
                });
            }
        })
        .catch(error => console.error('Error fetching data:', error));        
});
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
            allExperienceData = data.exps;
            displayExperience(); // 페이지 로드 시 초기 두 개의 경험만 표시
            
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
                                            <div class="text-primary fw-bolder mb-2">${experience.Period}</div>
                                            <div class="small fw-bolder">${experience.Role}</div>
                                            <div class="small text-muted">${experience.Company}</div>
                                            <div class="small text-muted">${experience.Location}</div>
                                        </div>
                                    </div>
                                    <div class="col-lg-8">
                                        <div>${experience.Description}</div>
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
            loadMoreButton.addEventListener('click', function () {
                displayExperience(); // 버튼 클릭 시 추가 데이터 로드
            });
            // Initial display of experiences (first call)
            //displayExperience();


        })
        .catch(error => console.error('Error fetching data:', error));

    // "더보기" 버튼 클릭 이벤트

});
