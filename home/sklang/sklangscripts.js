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
                        <div class="card shadow border-0 rounded-4 mb-5" data-id="${experience.ID}">
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
                            <button class="btn btn-light position-absolute top-0 start-0 m-2 p-1 edit-button" 
                            style="border-radius: 50%; z-index: 10;">
                            <i class="bi bi-pencil"></i>
                            </button>
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

        document.getElementById('uploadForm').addEventListener('submit', function (event) {
            event.preventDefault(); // 폼 제출 기본 동작 방지

            const period = document.getElementById('period').value;
            const role = document.getElementById('role').value;
            const company = document.getElementById('company').value;
            const description = document.getElementById('description').value;

            fetch('/uploadresume', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    period,
                    role,
                    company,
                    description
                })
            })
            .then(response => response.json())
            .then(result => {
                alert('Experience added successfully!');
                document.getElementById('newExperienceCard').remove(); // 새 카드 제거
            })
            .catch(error => console.error('Error:', error));
        });

        document.getElementById('cancelUpload').addEventListener('click', function () {
            document.getElementById('newExperienceCard').remove(); // 새 카드 제거
        });
    });

    // 수정 버튼 클릭 시 카드 내에 수정 폼 추가하기
    experienceContainer.addEventListener('click', function (event) {
        if (event.target.closest('.edit-button')) {
            const card = event.target.closest('.card');
            const cardId = card.getAttribute('data-id');

            // 카드에 수정 폼 추가
            const editCard = `
                <div class="card shadow border-0 rounded-4 mb-5" id="editExperienceCard">
                    <div class="card-body p-5">
                        <h3 class="fw-bolder mb-4">Edit Experience</h3>
                        <form id="editForm">
                            <input type="hidden" id="editId" value="${cardId}">
                            <div class="mb-3">
                                <label for="editPeriod" class="form-label">Period</label>
                                <input type="text" class="form-control" id="editPeriod" value="${card.querySelector('.text-primary').textContent}" required>
                            </div>
                            <div class="mb-3">
                                <label for="editRole" class="form-label">Role</label>
                                <input type="text" class="form-control" id="editRole" value="${card.querySelector('.small.fw-bolder').textContent}" required>
                            </div>
                            <div class="mb-3">
                                <label for="editCompany" class="form-label">Company</label>
                                <input type="text" class="form-control" id="editCompany" value="${card.querySelector('.small.text-muted').textContent}" required>
                            </div>
                            <div class="mb-3">
                                <label for="editDescription" class="form-label">Description</label>
                                <textarea class="form-control" id="editDescription" rows="3" required>${card.querySelector('.col-lg-8 div').textContent}</textarea>
                            </div>
                            <button type="submit" class="btn btn-primary">Save Changes</button>
                            <button type="button" id="cancelEdit" class="btn btn-secondary">Cancel</button>
                        </form>
                    </div>
                </div>
            `;

            card.insertAdjacentHTML('beforeend', editCard);

            document.getElementById('editForm').addEventListener('submit', function (event) {
                event.preventDefault(); // 폼 제출 기본 동작 방지

                const period = document.getElementById('editPeriod').value;
                const role = document.getElementById('editRole').value;
                const company = document.getElementById('editCompany').value;
                const description = document.getElementById('editDescription').value;
                const editId = document.getElementById('editId').value;

                fetch('/editresume', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        id: editId,
                        period,
                        role,
                        company,
                        description
                    })
                })
                .then(response => response.json())
                .then(result => {
                    alert('Experience updated successfully!');
                    document.getElementById('editExperienceCard').remove(); // 수정 카드 제거
                    // 여기에서 카드의 내용 업데이트 처리 필요
                })
                .catch(error => console.error('Error:', error));
            });

            document.getElementById('cancelEdit').addEventListener('click', function () {
                document.getElementById('editExperienceCard').remove(); // 수정 카드 제거
            });
        }
    });
});
