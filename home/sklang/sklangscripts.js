document.addEventListener("DOMContentLoaded", function () {
    // 서버에서 스킬 및 언어 데이터를 가져옴
    fetch('/returnskillang')
        .then(response => response.json())
        .then(data => {
            const skills = data.skills;
            const container = document.getElementById('skillsContainer');
            const languages = data.languages;
            const containerlang = document.getElementById('languagesContainer');   

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
    const experiencesPerPage = 2;

    fetch('/returnresume')
        .then(response => response.json())
        .then(data => {
            allExperienceData = data.exps;
            displayExperience();

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
                                        </div>
                                    </div>
                                    <div class="col-lg-8">
                                        <div>${experience.description}</div>
                                    </div>
                                </div>
                            </div>
                            <div>
                                <button class="btn btn-light position-absolute top-0 start-0 m-2 p-1 edit-button" 
                                style="border-radius: 50%; z-index: 10;"
                                data-id="${experience.ID}">
                                <i class="bi bi-pencil"></i>
                                </button>
                            </div>
                        </div>
                    `;

                    experienceContainer.insertAdjacentHTML('beforeend', experienceCard);
                }

                currentIndex = endIndex;

                if (currentIndex >= allExperienceData.length) {
                    loadMoreButton.style.display = 'none';
                }
            }

            loadMoreButton.addEventListener('click', function () {
                displayExperience();
            });
        })
        .catch(error => console.error('Error fetching data:', error));

    // 수정 버튼 클릭 시 모달 열기
    document.addEventListener('click', function (event) {
        if (event.target.closest('.edit-button')) {
            const button = event.target.closest('.edit-button');
            const experienceId = button.getAttribute('data-id');
            const experience = allExperienceData.find(exp => exp.ID === experienceId);

            if (experience) {
                // 모달에 데이터 채우기
                document.getElementById('editPeriod').value = experience.period;
                document.getElementById('editRole').value = experience.role;
                document.getElementById('editCompany').value = experience.company;
                document.getElementById('editLocation').value = experience.location;
                document.getElementById('editDescription').value = experience.description;

                // 모달 열기
                const editModal = new bootstrap.Modal(document.getElementById('editModal'));
                editModal.show();
                
                // 폼 제출 처리
                document.getElementById('editForm').addEventListener('submit', function (event) {
                    event.preventDefault(); // 폼 제출 기본 동작 방지

                    const updatedPeriod = document.getElementById('editPeriod').value;
                    const updatedRole = document.getElementById('editRole').value;
                    const updatedCompany = document.getElementById('editCompany').value;
                    const updatedLocation = document.getElementById('editLocation').value;
                    const updatedDescription = document.getElementById('editDescription').value;

                    // 수정된 데이터 서버에 전송
                    fetch('/updateexperience', {
                        method: 'POST',
                        headers: {
                            'Content-Type': 'application/json'
                        },
                        body: JSON.stringify({
                            id: experienceId,
                            period: updatedPeriod,
                            role: updatedRole,
                            company: updatedCompany,
                            location: updatedLocation,
                            description: updatedDescription
                        })
                    })
                    .then(response => response.json())
                    .then(result => {
                        console.log('Success:', result);
                        alert('Experience updated successfully!');
                        // 모달 닫기
                        editModal.hide();
                        // 화면 갱신
                        experienceContainer.innerHTML = ''; // 기존 카드 제거
                        displayExperience(); // 갱신된 카드 표시
                    })
                    .catch(error => console.error('Error:', error));
                });
            }
        }
    });
});
