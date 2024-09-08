document.addEventListener("DOMContentLoaded", function () {
    // 서버에서 스킬 및 언어 데이터를 가져옴
    fetch('/returnskillang')
        .then(response => response.json())
        .then(data => {
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
        })
        .catch(error => console.error('Error fetching data:', error));

    // 업로드 버튼 클릭 시 새로운 카드 열기
    document.getElementById('upload').addEventListener('click', function () {
        // 새 카드 추가
        const newCard = `
            <div class="card shadow border-0 rounded-4 mb-5" id="newExperienceCard">
                <div class="card-body p-5">
                    <h3 class="fw-bolder mb-4">Add New Experience</h3>
                    <form id="uploadForm">
                        <div class="mb-3">
                            <label for="period" class="form-label">Period</label>
                            <input type="text" class="form-control" id="period" required>
                        </div>
                        <div class="mb-3">
                            <label for="role" class="form-label">Role</label>
                            <input type="text" class="form-control" id="role" required>
                        </div>
                        <div class="mb-3">
                            <label for="company" class="form-label">Company</label>
                            <input type="text" class="form-control" id="company" required>
                        </div>
                        <div class="mb-3">
                            <label for="location" class="form-label">Location</label>
                            <input type="text" class="form-control" id="location" required>
                        </div>
                        <div class="mb-3">
                            <label for="description" class="form-label">Description</label>
                            <textarea class="form-control" id="description" rows="3" required></textarea>
                        </div>
                        <button type="submit" class="btn btn-primary">Submit</button>
                        <button type="button" id="cancelUpload" class="btn btn-secondary">Cancel</button>
                    </form>
                </div>
            </div>
        `;

        experienceContainer.insertAdjacentHTML('beforeend', newCard);

        // 폼 제출 처리
        document.getElementById('uploadForm').addEventListener('submit', function (event) {
            event.preventDefault(); // 폼 제출 기본 동작 방지

            const period = document.getElementById('period').value;
            const role = document.getElementById('role').value;
            const company = document.getElementById('company').value;
            const location = document.getElementById('location').value;
            const description = document.getElementById('description').value;

            // 폼 데이터 서버에 전송
            fetch('/uploadresume', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    period,
                    role,
                    company,
                    location,
                    description
                })
            })
            .then(response => response.json())
            .then(result => {
                console.log('Success:', result);
                // 업로드 성공 후 카드 닫기 및 피드백 표시
                alert('Experience added successfully!');
                document.getElementById('newExperienceCard').remove(); // 새 카드 제거
            })
            .catch(error => console.error('Error:', error));
        });

        // 취소 버튼 클릭 시 카드 제거
        document.getElementById('cancelUpload').addEventListener('click', function () {
            document.getElementById('newExperienceCard').remove(); // 새 카드 제거
        });
    });
});
